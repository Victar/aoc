package adventofcode;

import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;

public class DayX extends BaseTest {

	public static final int DAY = 0;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		for (final String input : data) {
			System.out.println(input);
		}
	}

}
