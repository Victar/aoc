package adventofcode;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;

// Util class to work with Graphs
// calculatePath - calculate min paths based on Dijkstra's algorithm

public class GraphUtil {

	public static void calculatePath(final Node from, final int initValue) {
		from.setDistance(initValue);
		final Set<Node> visited = new HashSet<>();
		final Set<Node> notVisited = new HashSet<>();
		notVisited.add(from);
		while (!notVisited.isEmpty()) {
			final Node current = getMinRiskNode(notVisited);
			notVisited.remove(current);
			for (final Map.Entry<Node, Integer> neighborEntry : current.getNeighbors().entrySet()) {
				final Node neighbor = neighborEntry.getKey();
				final Integer risk = neighborEntry.getValue();
				if (!visited.contains(neighbor)) {
					calculateMinRisk(neighbor, risk, current);
					notVisited.add(neighbor);
				}
			}
			visited.add(current);
		}
	}

	private static Node getMinRiskNode(final Set<Node> notVisited) {
		Node minRiskNode = null;
		int minRiskDistance = Integer.MAX_VALUE;
		for (final Node node : notVisited) {
			final int nodeDistance = node.getDistance();
			if (nodeDistance < minRiskDistance) {
				minRiskDistance = nodeDistance;
				minRiskNode = node;
			}
		}
		return minRiskNode;
	}

	private static void calculateMinRisk(final Node to, final Integer risks, final Node from) {
		final Integer fromDistance = from.getDistance();
		if (fromDistance + risks < to.getDistance()) {
			to.setDistance(fromDistance + risks);
			final LinkedList<Node> minPath = new LinkedList<>(from.getMinPath());
			minPath.add(from);
			to.setMinPath(minPath);
		}
	}

	public interface Node {

		int getDistance();

		void setDistance(final int distance);

		Map<Node, Integer> getNeighbors();

		List<Node> getMinPath();

		void setMinPath(List<Node> minPath);

	}
}
